/* http://keith-wood.name/countdown.html
   Dutch initialisation for the jQuery countdown extension
   Written by Mathias Bynens <http://mathiasbynens.be/> Mar 2008. */
(function($) {
	$.countdown.regional['nl'] = {
		labels: ['Jaren', 'Maanden', 'Weken', 'Dagen', 'Uren', 'Minuten', 'Seconden'],
		labels1: ['Jaar', 'Maand', 'Week', 'Dag', 'Uur', 'Minuut', 'Seconde'],
		compactLabels: ['j', 'm', 'w', 'd'],
		whichLabels: null,
		digits: ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'],
		timeSeparator: ':', isRTL: false};
	$.countdown.setDefaults($.countdown.regional['nl']);
})(jQuery);
